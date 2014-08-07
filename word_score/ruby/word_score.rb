class WordScore
  def self.score(word, pos=3)
    score = 0
    word.upcase.bytes.each do |b|
      break if pos < 0
      ns = b*(256**pos)
      score += ns
      pos -= 1
    end
    score
  end
end
