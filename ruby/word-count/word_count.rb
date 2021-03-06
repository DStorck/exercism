class Phrase

  def initialize(string)
    @words = string.downcase.scan(/\b[\w']+\b/)
  end

  def word_count
    counts = Hash.new(0)
    @words.each do | word |
      counts[word] += 1
    end
    return counts
  end
  
end

module BookKeeping
  VERSION = 1
end
