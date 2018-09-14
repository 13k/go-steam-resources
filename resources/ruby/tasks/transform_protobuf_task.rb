# frozen_string_literal: true

require 'pathname'

require_relative '../logging'
require_relative 'generate_file_task'

class TransformProtobufTask < GenerateFileTask
  include Logging

  attr_accessor :go_package

  def run
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
      say_status(:patch, "#{input_file_from_root} (#{patch.relative_path_from(root_path)})", :yellow)
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
