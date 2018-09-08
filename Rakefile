# frozen_string_literal: true

require 'pathname'
require 'tempfile'
require 'thread'
require 'yaml'

PARALLEL = ENV.fetch('PARALLEL', '1') == '1'

ROOT_PATH = Pathname.new(__dir__)
PROTOBUF_PKG_PATH = ROOT_PATH.join('protobuf')

BUILD_PATH = ROOT_PATH.join('build')
PATCHED_PROTO_PATH = BUILD_PATH.join('protobufs')
GENERATED_GO_PATH = BUILD_PATH.join('go')

RESOURCES_PATH = ROOT_PATH.join('resources')
CONFIG_PATH = RESOURCES_PATH.join('config.yml')
PATCHES_PATH = RESOURCES_PATH.join('patches')

LOG_M = Mutex.new
COLORS = {
  reset:    0,
  black:   30,
  red:     31,
  green:   32,
  yellow:  33,
  blue:    34,
  magenta: 35,
  cyan:    36,
  white:   37,
}.freeze

# Hack to synchronize rake output messages.
# This won't avoid out-of-order lines from multiple tasks, it will only avoid a line being written
# before a previous line finished. It works for single-line status messages.
module Rake
  class << self
    alias_method :rake_output_message_orig, :rake_output_message
    def rake_output_message(*args, &block)
      LOG_M.synchronize do
        rake_output_message_orig(*args, &block)
        $stderr.flush
      end
    end
  end
end

def colorize(text, color_code)
  color_code = COLORS.fetch(color_code) if color_code.is_a?(Symbol)
  "\033[#{color_code}m#{text}\033[0m"
end

COLORS.each do |key, color_code|
  define_method(key) do |text|
    colorize(text, color_code)
  end
end

def say(msg)
  LOG_M.synchronize do
    $stderr.puts(msg)
    $stderr.flush
  end
end

def say_status(status, message, color = :reset)
  say("#{colorize(status.to_s.rjust(11), color)} #{colorize(message, :reset)}")
end

def symbolize_keys(value)
  case value
  when Hash
    value
      .transform_keys { |key| key.to_sym }
      .transform_values { |val| symbolize_keys(val) }
  when Enumerable
    value.map { |item| symbolize_keys(item) }
  else
    value
  end
end

class ProtobufTask < Rake::FileTask
  include Rake::FileUtilsExt

  def initialize(*args, &block)
    super
    @executed = false
  end

  def executed?
    @executed
  end

  def executed!
    @executed = true
  end

  protected

  def output_file
    @output_file ||= Pathname.new(name)
  end

  def output_file_from_root
    @output_file_from_root ||= output_file.relative_path_from(ROOT_PATH)
  end

  def output_dir
    @output_dir ||= output_file.dirname
  end

  def input_files
    @input_files ||= prerequisite_tasks.map(&:name).map { |path| Pathname.new(path) }
  end

  def input_files_from_root
    @input_files_from_root ||= input_files.map { |p| p.relative_path_from(ROOT_PATH) }
  end

  def temp_file
    @temp_files ||= []

    Tempfile.new.tap do |temp_file|
      @temp_files << temp_file
    end
  end

  def clean_temp_files
    return unless @temp_files
    @temp_files.each(&:unlink)
  end
end

class GenerateProtobufTask < ProtobufTask
  attr_accessor :include_paths, :paths, :plugins, :imports, :import_path, :output_path

  def initialize(*args, &block)
    super

    @include_paths = []
    @plugins = []
    @imports = {}
  end

  def generate
    return if executed?

    executed!
    say_status(:generate, output_file_from_root, :green)
    compile_file
    install_file
  end

  protected

  def protoc_options
    [
      *proto_path_options,
      *go_out_option,
    ]
  end

  def proto_path_options
    paths = @include_paths

    if paths.empty?
      paths += input_files.map(&:dirname)
    end

    paths.flat_map { |path| ['-I', path.to_s] }
  end

  def go_out_option
    options = {
      paths: @paths,
      plugins: go_out_plugins,
      **go_out_imports,
    }
      .reject { |_, v| v.nil? }
      .map { |k, v| "#{k}=#{v}" }
      .join(',')

    [format('--go_out=%<options>s:%<output>s', options: options, output: @output_path)]
  end

  def go_out_plugins
    @plugins.empty? ? nil : @plugins.join('+')
  end

  def go_out_imports
    @imports.transform_keys { |proto| "M#{proto}" }
  end

  def compile_file
    mkdir_p(@output_path)
    sh('protoc', *protoc_options, *input_files.map(&:to_s))
  end

  def install_file
    mkdir_p(output_dir)
    src = @output_path.join(@import_path, output_file.basename)
    install(src, output_file)
  end
end

class TransformProtobufTask < ProtobufTask
  attr_accessor :go_package

  def transform
    return if executed?

    executed!
    mkdir_p(output_dir)
    transformed = apply_patches(input_file)
    transformed = inject_fixes(transformed)
    save(transformed)
  ensure
    clean_temp_files
  end

  def patches_path=(path)
    @patches_path = Pathname.new(path)
  end

  protected

  def input_file
    @input_file ||= input_files.first
  end

  def input_file_from_root
    @input_file_from_root ||= input_files_from_root.first
  end

  def patches
    @patches ||= begin
      pattern = input_file.basename.sub_ext('.proto-[0-9][0-9].patch')
      @patches_path.glob(pattern).sort
    end
  end

  def patch_options(patch, output)
    [
      '--batch',
      '--quiet',
      '-i', patch.to_s,
      '-o', output.to_s,
    ]
  end

  def apply_patches(input)
    return input if patches.empty?
    output = nil

    patches.each do |patch|
      output = Pathname.new(temp_file)
      say_status(:patch, "#{input_file_from_root} (#{patch.relative_path_from(ROOT_PATH)})", :yellow)
      sh('patch', *patch_options(patch, output), input.to_s)
      input = output
    end

    output
  end

  def inject_fixes(input)
    output_file = temp_file

    sed_cmd = <<~SED
      1 i\\
      syntax = "proto2";\\
      option go_package = "#{@go_package}";\\
    SED

    say_status(:fix, "#{input_file_from_root}", :cyan)
    sh('sed', '-e', sed_cmd, input.to_s, out: output_file)
    Pathname.new(output_file)
  end

  def save(input)
    cp(input, output_file)
  end
end

def generate_protobuf(*args, &block)
  task = GenerateProtobufTask.define_task(*args, &block)
  task.enhance { |t, _args| t.generate }
end

def transform_protobuf(*args, &block)
  task = TransformProtobufTask.define_task(*args, &block)
  task.enhance { |t, _args| t.transform }
end

config = symbolize_keys(YAML.load_file(CONFIG_PATH))
proto_path = ROOT_PATH.join(config[:input_rel_path])

tasklist = Hash.new { |h, k| h[k] = [] }
config[:packages].each_with_object(tasklist) do |package, tasks|
  pkg_rel_path = Pathname.new(package[:name])
  pkg_import_path = File.join(config[:base_import_path], package[:name])
  pkg_patches_path = PATCHES_PATH.join(pkg_rel_path)
  pkg_patched_path = PATCHED_PROTO_PATH.join(pkg_rel_path)
  pkg_output_path = PROTOBUF_PKG_PATH.join(pkg_rel_path)

  excludes = package
    .fetch(:exclude_files, [])
    .map { |pattern| proto_path.join(pattern).to_s }

  package[:files].flat_map do |pattern|
    proto_path
      .glob(pattern)
      .reject { |file| excludes.any? { |excl_pattern| file.fnmatch?(excl_pattern) } }
      .map do |input_file|
        transformed_file = pkg_patched_path.join(input_file.basename)

        transform_task = transform_protobuf(transformed_file => input_file) do |task|
          task.patches_path = pkg_patches_path
          task.go_package = pkg_import_path
        end

        tasks[:transform] << transform_task

        output_file = pkg_output_path.join(input_file.basename.sub_ext('.pb.go'))

        tasks[:generate] << generate_protobuf(output_file => transform_task) do |task|
          task.import_path = pkg_import_path
          task.output_path = GENERATED_GO_PATH
        end
      end
  end
end

task_def_method = PARALLEL ? :multitask : :task
send(task_def_method, transform: tasklist[:transform])
send(task_def_method, generate: tasklist[:generate])

task default: %i[transform generate]

task :clean do
  rm_rf BUILD_PATH
end

task :clean_proto do
  rm_rf ROOT_PATH.join('protobuf')
end

task distclean: %i[clean clean_proto]
