# frozen_string_literal: true

require_relative '../logging'
require_relative 'generate_file_task'

class GenerateSteamLangTask < GenerateFileTask
  include Logging

  attr_accessor :gen_go_file, :go_package, :protobuf_package

  def run
    say_status(:steamlang, output_file_from_root, :green)
    generate_file
  end

  protected

  def input_file
    @input_file ||= input_files.first
  end

  def gen_go_options
    [
      '-pkg', @go_package,
      '-protopkg', @protobuf_package,
      '-o', output_file.to_s,
    ]
  end

  def generate_file
    mkdir_p(output_dir)
    sh('go', 'run', @gen_go_file.to_s, *gen_go_options, input_file.to_s)
  end
end
