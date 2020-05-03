# frozen_string_literal: true

require_relative 'logging'
require_relative 'generate_file_task'

class GenerateProtobufTask < GenerateFileTask
  include Logging

  attr_accessor :include_path, :output_path, :go_opt

  def initialize(*args, &block)
    super

    @include_path = nil
    @output_path = nil
    @go_opt = {}
  end

  # TODO: remove input file if failed?
  def run
    require_command!('protoc')
    require_command!('protoc-gen-go')

    say_status(:protoc, output_file_from_root, :green)

    compile_file
    install_file
  end

  protected

  def input_file
    @input_file ||= input_files.first
  end

  def input_file_from_root
    @input_file_from_root ||= input_files_from_root.first
  end

  def protoc_options
    @protoc_options ||= [
      *proto_path_options,
      *go_out_option.to_a,
      *go_opt_option.to_a,
    ]
  end

  def proto_path_options
    raise Rake::TaskArgumentError, 'include_path cannot be empty' if @include_path.nil?

    ['-I', @include_path.to_s]
  end

  def go_out_option
    raise Rake::TaskArgumentError, 'output_path cannot be empty' if @output_path.nil?

    [format('--go_out=%<go_out>s', go_out: @output_path)]
  end

  def go_opt_option
    return if @go_opt.empty?

    go_opt_str = @go_opt
      .map { |k, v| "#{k}=#{v}" }
      .join(',')

    [format('--go_opt=%<go_opt_str>s', go_opt_str: go_opt_str)]
  end

  def generated_file
    @generated_file ||= begin
        relpath = input_file.relative_path_from(@include_path)
        @output_path.join(relpath.sub_ext('.pb.go'))
      end
  end

  def compile_file
    mkdir_p(@output_path)
    sh('protoc', *protoc_options, input_file.to_s)
  end

  def install_file
    mkdir_p(output_dir)
    install(generated_file, output_file)
  end
end

def generate_protobuf(*args, &block)
  GenerateProtobufTask.define_task(*args, &block)
end
