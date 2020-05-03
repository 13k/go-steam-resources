# frozen_string_literal: true

require 'yaml'

module Util
  CommandNotFoundError = Class.new(RuntimeError)

  def self.go_module(module_file)
    which!('go')

    Dir.chdir(module_file.dirname) do
      `go list -m`.strip
    end
  end

  def self.load_config(config_file)
    YAML.safe_load(config_file.read, symbolize_names: true)
  end

  def self.which(cmd)
    system('which', cmd, out: :close, err: :close)
  end

  def self.which!(cmd)
    return if which(cmd)

    raise CommandNotFoundError, format('command %<cmd>s not found', cmd: cmd)
  end
end
