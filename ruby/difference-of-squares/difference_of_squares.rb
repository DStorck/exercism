class Squares

  def initialize(num)
    @num = num
  end

  def square_of_sum
    nums = (1..@num).to_a
    sum = nums.reduce(:+)
    return sum ** 2
  end

  def sum_of_squares
    total = 0
    @num.times do | num |
      total += (num + 1)  ** 2
    end
    return total
  end

  def difference
    return square_of_sum - sum_of_squares
  end

end

module BookKeeping
  VERSION = 4
end
