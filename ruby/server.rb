require 'socket'

server = TCPServer.new 80
pids = []

loop do
  Thread.start(server.accept) do |client|
    sleep 3600
    # headers = []
    # while header = client.gets
    #   break if header.chomp.empty?
    #   headers << header.chomp
    # end

    # puts Time.new.to_s + headers.join(", ")

    # client.puts "HTTP/1.0 200 OK"
    # client.puts "Content-Type: text/plain"
    # client.puts
    # client.puts Time.new.to_s + headers.join(", ")
    client.close
  end
end
