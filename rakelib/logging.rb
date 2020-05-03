# frozen_string_literal: true

require 'rake'

LOG_M = Mutex.new
COLORS = {
  reset: 0,
  black: 30,
  red: 31,
  green: 32,
  yellow: 33,
  blue: 34,
  magenta: 35,
  cyan: 36,
  white: 37,
}.freeze

# Hack to synchronize rake output messages.
# This won't avoid out-of-order lines from multiple tasks, it will only avoid a line being written
# before a previous line finished. It works for single-line status messages.
module Rake
  class << self
    alias rake_output_message_orig rake_output_message

    def rake_output_message(*args, &block)
      LOG_M.synchronize do
        rake_output_message_orig(*args, &block)
        $stderr.flush
      end
    end
  end
end

module Logging
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
      $stderr.puts(msg) # rubocop:disable Style/StderrPuts
      $stderr.flush
    end
  end

  def say_status(status, message, color = :reset)
    say("#{colorize(status.to_s.rjust(11), color)} #{colorize(message, :reset)}")
  end
end
