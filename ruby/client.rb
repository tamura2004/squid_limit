sleep 3
60000.times do |n|
	5.times do |m|
		Thread.new do
			`curl web#{m+1} -x sq:3128`
		end
	end
end


# require "socket"

# def get
#   TCPSocket.open("sq", 3128) do |sock|
#     sock.print "GET http://web/ HTTP/1.0\r\n\r\n"
#     puts sock.read
#   end
# end

# 60000.times do |n|
#   Thread.new do
#     # puts "number: #{n}"
#     get
#   end
# end

# sleep 3600
