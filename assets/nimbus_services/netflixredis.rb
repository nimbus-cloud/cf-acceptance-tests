require 'sinatra/base'
require 'redis'

module Nimbus
  class ServicesApp < Sinatra::Base

    get '/netflixredis/insert/:key/:value' do
      netflixredis = Redis.new(:url => Nimbus::Config.netflixredis)
      netflixredis.set(params['key'], params['value'])
      'OK'
    end

    get '/netflixredis/read/:key/:value' do
      netflixredis = Redis.new(:url => Nimbus::Config.netflixredis)
      netflixredis.get(params['key']) == params['value'] ? 'OK' : 'FAIL'
    end

  end
end
