
@board = [1, 2, 3, 4, 5, 6, 7, 8, 9]
@player1 = []
@player2 = []


def p1_turn
  def move(n)
    if @board.include?(n)
      @player1.push(n)
      @board.delete(n)
    else
      puts "Play not available, please try again!"
      p1_turn
    end
  end
  puts "Player 1, what have you?"
  choice = gets.chomp.to_i
  move(choice)
end

def p2_turn
  def move(n)
    if @board.include?(n)
      @player2.push(n)
      @board.delete(n)
    else
      puts "Play not available, please try again!"
      p2_turn
    end
  end
  puts "Player 2, what have you?"
  choice = gets.chomp.to_i
  move(choice)
end


  puts "Board starts with: #{@board}"
  puts "Player 1 starts with: #{@player1}"
  puts "Player 2 starts with: #{@player2}"
  puts

p1_turn

  puts
  puts "End of turn 1a:"
  puts "Board now has available: #{@board}"
  puts "Player 1 now has: #{@player1}"
  puts "Player 2 now has: #{@player2}"
  puts

p2_turn

  puts
  puts "End of turn 1b:"
  puts "Board now has available: #{@board}"
  puts "Player 1 now has: #{@player1}"
  puts "Player 2 now has: #{@player2}"
  puts
