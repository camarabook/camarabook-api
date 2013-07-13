require "rubygems"
require "bundler"
Bundler.require

StandaloneMigrations::Tasks.load_tasks

namespace :data do
  task :enviroment do
    ENV['USE_PARSER'] = 'true'
    require './config/boot'
  end

  desc "get cota data"
  task :cotas => :enviroment do
    CotaCrawlerService.save_from_cota_xml_parser
  end

  desc "get deputados"
  task :deputados => :enviroment do
    DeputadoCrawlerService.save_from_pesquisa_parser
    DeputadoCrawlerService.save_from_deputado_xml_parser
  end

  desc "run all data tasks"
  task :all => [:deputados, :cotas]
end
