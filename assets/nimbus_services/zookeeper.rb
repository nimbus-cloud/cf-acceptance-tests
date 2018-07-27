require 'sinatra/base'
require 'zookeeper'

# scms-zookeeper

module Nimbus
	class ServicesApp < Sinatra::Base 

		configure do 
			z = Zookeeper.new(Nimbus::Config.zookeeper['cluster-nodes'] + ":2181")
			z.get_children(:path => "/")
		end


		get 'zookeeper/create/:key/:value' do |key, value|
			Zookeeper.add_auth(Nimbus::Config.zookeeper['user'], Nimbus::Config.zookeeper['password'])
			Zookeeper.create(:path => "/#{key}/#{value}")
		end

		get 'zookeeper/get/:key/:value' do |key, value|
			Zookeeper.add_auth(Nimbus::Config.zookeeper['user'], Nimbus::Config.zookeeper['password'])
			secret = Zookeeper.get(:path => "/#{key}")
			secret.data[key.to_sym] == value ? 'OK' : 'FAIL'
		end

		end
	end
end
