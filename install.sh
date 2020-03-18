# Purpose of this file is to facilitate easy installation of golang on your debian box
# Since most package managers have older versions of golang which may result in errors while running the tool
# use this instead to install golang and the tool!

if [ -z $GOPATH ]
then
  echo -e "\e[32m Installing golang \e[0m"
  wget https://dl.google.com/go/go1.14.src.tar.gz && tar -C /usr/local -xzf go1.14.src.tar.gz
  echo "export PATH=$PATH:/usr/local/go/bin:$GOPATH/bin" >> ~/.bashrc && source ~/.bashrc
  echo -e "\e[32m Installed golang. Test it out by running \"go version\" \e[0m"
fi

echo -e "\e[32m Building hakrevdns \e[0m"
go build && mv hakrevdns /usr/local

if [ -e /usr/bin/hakrevdns ]
then
  echo -e "\e[31mhakrevdns \e[32m has been successfully installed\e[0m"
  echo -e "\e[32m Test it out by trying, echo "173.0.84.110" | hakrevdns \e[0m"
fi


# can't thank hakluke enough for making these amazing, fast tools, they give results in no time, amazingly easing out the workflow
