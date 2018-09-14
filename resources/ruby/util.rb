# frozen_string_literal: true

require 'yaml'

def parse_go_module(module_file)
  module_file.readlines(chomp: true).grep(/^module/).first.sub(/^module\s*/, '')
end

def load_config(config_file)
  YAML.safe_load(config_file.read, symbolize_names: true)
end
