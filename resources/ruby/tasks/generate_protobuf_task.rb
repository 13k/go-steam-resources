# frozen_string_literal: true

require_relative '../logging'
require_relative 'generate_file_task'

class GenerateProtobufTask < GenerateFileTask
  include Logging

  attr_accessor :include_paths, :paths, :plugins, :imports, :import_path, :output_path

  def initialize(*args, &block)
    super

    @include_paths = []
    @plugins = []
    @imports = {}
  end

  # TODO: remove input file if failed?
  def run
    say_status(:protobuf, output_file_from_root, :green)
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
