Rails.application.routes.draw do
  namespace :metadata do
    get '/names/:name', to: 'store#get_name'
    post '/names/:name', to: 'store#incr_name'
  end
end
