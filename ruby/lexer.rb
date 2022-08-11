module Lexer
  OPERATORS = ["+", "-", "*", "/"]
  OPERATORS_REGEX = Regexp.new("([" + Regexp.quote(OPERATORS.join("")) + "])")
  TYPE_OPERATOR = "operator"
  TYPE_NUMBER = "number"

  class Token
    attr_accessor :type, :value

    def initialize(type, value)
      @type = type
      @value = value
    end

    def to_s
      @type + " " + @value
    end
  end

  def stream(statement)
    # puts "debug:lexer:" + statement
    stream = Array.new
    splitted = statement.split(OPERATORS_REGEX)
    splitted.each {|element|
      if element.chomp.empty?
      elsif OPERATORS.include?(element)
        stream.push(Token.new(TYPE_OPERATOR, element))
      else Float element
        stream.push(Token.new(TYPE_NUMBER, element))
      end
    }

    # stream.each {|element|
    #   puts "debug:lexer:element:" + element.to_s
    # }
    stream
  end

end