class Complement

CONVERSIONS=  {
   "G" => "C",
   "C" => "G",
   "T" => "A",
   "A" => "U"
}
  def self.of_dna(str)
    converted_string = ""
    chars = str.chars
    chars.each do | char |
      if CONVERSIONS[char]
        converted_string +=  CONVERSIONS[char]
      else
        return ""
      end
    end
    return converted_string
  end
end

module BookKeeping
  VERSION = 4
end


puts Complement.of_dna("G")
