# -*- mode: ruby -*-
# vi: set ft=ruby :

Vagrant.configure("2") do |config|

  config.vm.box     = "precise64"

  if Vagrant.has_plugin?("vagrant-cachier")
    config.cache.scope = :box
  end

  config.vm.provision "shell", inline: <<-SHELL
    if [ ! -f "go1.4.2.linux-amd64.tar.gz" ]; then
      wget https://storage.googleapis.com/golang/go1.4.2.linux-amd64.tar.gz
      tar -C /usr/local -xzf go1.4.2.linux-amd64.tar.gz
    fi
    apt-get -y install git-core
    if [ ! -d "dotfiles" ]; then
      git clone https://github.com/urbanautomaton/dotfiles
    fi
    chown vagrant:vagrant -R .
    sudo -i -u vagrant sh -c 'cd dotfiles && ./install.sh'
  SHELL

  # config.vm.define :prof do |vm_config|
  #   vm_config.vm.network :private_network, :ip => ip_address
  #   vm_config.vm.provision :shell do |s|
  #     s.path = "script/bootstrap"
  #   end
  # end

end

