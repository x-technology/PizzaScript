require_relative "lexer.rb"
require_relative "parser.rb"
require_relative "interpreter.rb"
include Lexer
include Parser
include Interpreter

puts "Welcome to Calculator!"
statement = nil

def exec(statement)
  puts "debug:exec:" + statement
  stream = Lexer::stream(statement)
  tree = Parser::tree(stream)
  puts Interpreter::eval(tree.root)
end

exec("1+2")

while statement != "q"
  puts "what should I do for you?"
  statement = gets.chomp
  exec(statement)
end

puts "just lucky"
