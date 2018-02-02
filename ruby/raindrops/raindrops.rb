class Raindrops

  def self.convert(int)
    rain = ""
    if int % 3 == 0
      rain += "Pling"
    end
    if int % 5 == 0
      rain += "Plang"
    end
    if int % 7 == 0
      rain += "Plong"
    end
    if rain == ""
      return int.to_s
    end
    return rain
  end

end

module BookKeeping
  VERSION = 3
end
