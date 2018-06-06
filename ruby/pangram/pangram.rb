class Pangram

  def self.pangram?( string )
    string = string.gsub(/[^A-Za-z]/, '')
    string.downcase!
    return false if string.length < 26
    letters = {}
    string.chars.each do |letter|
      letters[letter] = true
    end
    return letters.length == 26 ? true : false
  end


end

module BookKeeping
  VERSION = 6
end
