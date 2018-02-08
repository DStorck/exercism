class Squares

  def initialize(num)
    @num = num
  end

  def square_of_sum
    (1..@num).reduce(:+) ** 2
  end

  def sum_of_squares
    squares = (1..@num).collect { |x| x ** 2 }
    squares.reduce(:+)
  end

  def difference
    return square_of_sum - sum_of_squares
  end

end

module BookKeeping
  VERSION = 4
end
