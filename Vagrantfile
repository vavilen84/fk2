# -*- mode: ruby -*-
# vi: set ft=ruby :
# All Vagrant configuration is done here. For a complete reference,
# please see the online documentation at vagrantup.com.

VAGRANTFILE_API_VERSION = "2"
Vagrant.require_version ">= 1.7.2"

#detect envirinmnet variables
host_os = RbConfig::CONFIG['host_os']
vagrant_command = ARGV[0]

if vagrant_command == "up"
    # check and install required Vagrant plugins
	required_plugins = %w( vagrant-vbguest vagrant-cachier vagrant-hostsupdater vagrant-host-shell )
	required_plugins.each do |plugin|
		unless Vagrant.has_plugin?(plugin)
			system "vagrant plugin install #{plugin}" unless Vagrant.has_plugin? plugin
		end
	end
	# install nfs-kernel-server in Linux to support synced_folder via nfs
    if host_os =~ /linux/
    	system "sudo apt-get install nfs-kernel-server nfs-common rpcbind"
    end
end

Vagrant.configure(VAGRANTFILE_API_VERSION) do |config|

  config.vm.boot_timeout = 300
  config.vm.provider "virtualbox" do |vb|
  	vb.gui = false
  end

  config.vm.define "fotokolo" do |company|
	company.vm.box = "fotokolo"
	# Configure some Virtual Box params
	company.vm.provider :virtualbox do |company|
		company.customize ["modifyvm", :id, "--name", "fk2"]
		company.customize ["modifyvm", :id, "--ostype", "Ubuntu_64"]
		company.customize ["modifyvm", :id, "--memory", "1500"]
		company.customize ["modifyvm", :id, "--cpuexecutioncap", "80"]
		company.customize ["modifyvm", :id, "--natdnshostresolver1", "on"]
        company.customize ["modifyvm", :id, "--natdnsproxy1", "on"]
		# Set VirtualBox guest CPU count to the number of host cores
		if host_os =~ /linux/
        	company.customize ["modifyvm", :id, "--cpus", `grep "^processor" /proc/cpuinfo | wc -l`.chomp ]
        end
	end
	company.vm.box_url = "https://cloud-images.ubuntu.com/xenial/current/xenial-server-cloudimg-amd64-vagrant.box"
        company.vm.network "private_network", ip: "10.0.77.77"
        company.vm.network "forwarded_port", guest: 80, host: 8081

		#share synced_folder via nfs for Mac nd Linux
        if host_os =~ /linux/ or host_os =~ /darwin/
			company.vm.synced_folder "/var/www/fk2", "/var/www/fk2", type: "nfs", version: 3, nfs_udp: 1, mount_options: ["rw", "tcp", "nolock", "noacl", "sync"]
        else
			company.vm.synced_folder "/var/www/fk2", "/var/www/fk2", owner: "www-data", group: "www-data"
        end

        # PLUGINS
        # Set entries in hosts file
        # https://github.com/cogitatio/vagrant-hostsupdater
        if Vagrant.has_plugin?("vagrant-hostsupdater")
          company.hostsupdater.remove_on_suspend = true
          company.vm.hostname = "10.0.77.77"

          config.hostsupdater.aliases = [
            "db.local",
            "fotokolo.local"
            ]
        end
        if Vagrant.has_plugin?("vagrant-cachier")
          company.cache.auto_detect = true
        end

		# PROVISIONING
		# Ansible
		# To use Ansible provisioning you should have Ansible installed on your host machine
		# see here http://docs.ansible.com/intro_installation.html#installing-the-control-machine
		company.vm.provision "ansible" do |ansible|
			# should be equal to host name in Ansible hosts file
			ansible.limit = "dev"
			ansible.playbook = "provision/dev.yml"
			ansible.inventory_path = "provision/dev"
			# set to 'vvvv' for debug output in case of problems or leave it false
			ansible.verbose = 'vvvv'
			ansible.host_key_checking = false
		end

  end

end