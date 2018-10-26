require "redis"

class Metadata::StoreController < ApplicationController
  @@redis = Redis.new

  def get_name
    render json: { count: count_for_name }
  end

  def incr_name
    @@redis.incr params[:name]
    render json: { count: count_for_name }
  end

  private

  def count_for_name
    @@redis.get(params[:name]).to_i || 0
  end
end