# frozen_string_literal: true

require 'pathname'

require_relative 'logging'
require_relative 'protobuf'
require_relative 'generate_file_task'

# rubocop:disable Metrics/ClassLength

class TransformProtobufTask < GenerateFileTask
  include Logging

  IMPORT_RE = /^import\s+"(?<path>[^"]+)"\s*;\s*$/.freeze
  SELF_REF_RE = /(?<prefix>[^[:alnum:]])(?<ref>\.(?<sym>\w+))/.freeze

  def inputs=(inputs)
    @inputs = Array(inputs).map { |i| validate_input!(i) }
  end

  def run
    file = input_file
    file = apply_patches(file)
    file = transform(file)

    save_file(file)
  ensure
    clean_temp_files
  end

  protected

  def require_protobuf_file!(arg)
    unless arg.is_a?(ProtobufFile)
      raise ArgumentError, format(
        '%<task>s requires inputs to be ProtobufFile objects, got %<actual>s',
        task: self.class.name,
        actual: arg.class.name,
      )
    end

    arg
  end

  def validate_input!(arg)
    return arg if arg.is_a?(TransformProtobufTask)

    require_protobuf_file!(arg)
  end

  def input_file
    @input_file ||= input_files.first
  end

  def input_file_from_root
    @input_file_from_root ||= input_files_from_root.first
  end

  def package
    @package ||= input_file.package
  end

  def patches
    @patches ||= begin
        pattern = input_file.basename.sub_ext('.proto-[0-9][0-9].patch')
        package.patches_path.glob(pattern).sort
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

      say_status(
        :patch,
        "#{input_file_from_root} (#{patch.relative_path_from(root_path)})",
        :yellow,
      )

      sh('patch', *patch_options(patch, output), input.to_s)

      input = output
    end

    output
  end

  def transform(input)
    say_status(:transform, input_file_from_root, :cyan)

    tmp = temp_file

    begin
      tmp.write(header)
      tmp.write(transform_body(input))
    ensure
      tmp.close
    end

    Pathname.new(tmp)
  end

  def header
    @header ||= <<~HEADER
      syntax = "proto2";

      package #{package.pb_package};

      option go_package = "#{package.go_package}";

    HEADER
  end

  def transform_body(input)
    imports = []

    lines = input.each_line.map do |line|
      case line
      when IMPORT_RE
        transform_import_line(line, re_match: Regexp.last_match, imports: imports)
      when SELF_REF_RE
        transform_self_ref_line(line, imports: imports)
      else
        line
      end
    end

    lines.join
  end

  def transform_import_line(line, re_match:, imports:)
    path = re_match[:path]

    return line if path.start_with?('google/')

    file = resolve_import!(path).tap do |f|
      imports << f
    end

    path = File.join(file.package.path, path)

    format(%(import "%<path>s";\n), path: path)
  end

  def resolve_import(path)
    abspath = input_file.dirname.join(path)
    package.repo[abspath]
  end

  def resolve_import!(path)
    resolve_import(path).tap do |file|
      if file.nil?
        raise format(
          'could not find package for import %<path>s',
          path: path,
        )
      end
    end
  end

  def transform_self_ref_line(line, imports:)
    scope = [input_file, *imports]

    line.gsub(SELF_REF_RE) do
      match = Regexp.last_match
      ref = case match[:ref]
        when '.google'
          'google'
        else
          pkg = resolve_symbol!(match[:sym], scope, file: input_file)
          "#{pkg.pb_package}#{match[:ref]}"
        end

      "#{match[:prefix]}#{ref}"
    end
  end

  def resolve_symbol(sym, scope)
    scope.find { |f| f.symbols.find { |s| s[:id] == sym } }&.package
  end

  def resolve_symbol!(sym, scope, file:)
    resolve_symbol(sym, scope).tap do |pkg|
      if pkg.nil?
        raise format(
          'could not find package for symbol %<sym>s referenced in file %<file>s',
          sym: sym,
          file: file,
        )
      end
    end
  end

  def save_file(input)
    mkdir_p(output_dir)
    cp(input, output_file)
  end
end

# rubocop:enable Metrics/ClassLength

def transform_protobuf(output_file, proto_file, &block)
  TransformProtobufTask.define_task(output_file, proto_file, &block)
end
