require_relative "lexer.rb"
include Lexer

module Parser
=begin
  fun expr(rbp = 0) {
    // null denotation
    var left = nud(next())

    while bp(peek()) > rbp
      // left denotation
      left = led(left, next())
    left
  }

  fun nud(operator) {
    node(operator, expr())
  }
  fun led(left, operator) {
    node(left, operator, expr(bp(operator))
  }
=end
  PRECEDENCES = {
    "+" => 1,
    "*" => 2
  }

  class Node
    attr_accessor :left, :center, :right

    def to_s
      print "{"
      if !@left.nil?
        print @left.to_s
      end
      print ","
      if !@center.nil?
        print @center.to_s
      end
      print ","
      if !@right.nil?
        print @right.to_s
      end
      print "}"
    end
  end

  class Tree
    attr_accessor :stream, :root

    def initialize(stream)
      @stream = stream.dup

      @root = expr()
    end

    def bp(operator)
      # puts "debug:parser:bp:" + operator.value
      return PRECEDENCES[operator.value] || 0
    end

    def peek
      @stream[0]
    end

    def next
      if @stream.count == 0
        return
      end

      @stream.shift
    end

    def expr(rbp = 0)
      left = self.nud(self.next)
      # puts left.to_s
      while @stream.count > 0 and bp(peek()) > rbp
        left = self.led(left, self.next)
      end

      left
    end

    def led(element, operator)
      node = Node.new()
      node.left = element
      node.center = operator
      node.right = self.expr(bp(operator))
      node
    end

    def nud(element)
      if element.nil?
      elsif element.type == Lexer::TYPE_NUMBER
        node = Node.new()
        node.left = element
        # puts node.to_s
        return node
      elsif element.value == "-" or element.value == "+"
        # puts "debug:parser:nud:center:+-"
        node = Node.new()
        node.center = element
        node.left = self.expr
        return node
      end
      return nil
    end

    def to_s
      root.to_s
    end
  end

  def tree(stream)
    tree = Tree.new(stream)
    # puts tree.to_s
    tree
  end

end