module Interpreter

  def eval(node)
    # puts "debug:interpreter:eval"
    # puts node.to_s
    operator = node.center
    left = node.left
    right = node.right

    if operator.nil?
      return Float left.value
    end

    if right.nil?
      return -eval(left)
    end

    case operator.value
    when "+"
      return eval(right) + eval(left)
    when "*"
      return eval(left) * eval(right)
    end
  end

end