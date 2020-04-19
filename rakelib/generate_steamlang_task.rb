# frozen_string_literal: true

require_relative 'logging'
require_relative 'generate_file_task'

class GenerateSteamLangTask < GenerateFileTask
  include Logging

  attr_accessor :gen_package, :go_package, :protobuf_package

  def run
    require_command!('go')

    say_status(:steamlang, output_file_from_root, :green)

    generate_file
  end

  protected

  def input_file
    @input_file ||= input_files.first
  end

  def gen_options
    @gen_options ||= [
      '-pkg', @go_package,
      '-protopkg', @protobuf_package,
      '-o', output_file.to_s,
    ]
  end

  def generate_file
    mkdir_p(output_dir)

    sh('go', 'run', @gen_package.to_s, *gen_options, input_file.to_s)
  end
end

def generate_steamlang(*args, &block)
  GenerateSteamLangTask.define_task(*args, &block)
end
