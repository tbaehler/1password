# 1Password anywhere server for linux

You have 1Password, and sync it with Dropbox.
<br>Get the stupid simple go server for 1Password anywhere

Download it for any linux distro as binary from [here](https:///github.com/tbaehler/1password/raw/master/binary/1password).

Or install it from GO source

#####installation from GO source
<pre><code>go get github.com/tbaehler/1password
#start it from your GOPATH/bin
1password
</code></pre>

#####installation as binary
<pre><code>run the server
chmod +x 1password
>./1password
</code></pre>

if you do not have 1Password.html at the default location
$(HOME)/Dropbox/1Password.agilekeychain/
you can set an env variable with the directory that contains your 1Password.html

<pre><code>
>export ONEPASSWORD_DIR="your directory that holds the dropbox 1Password.html"
>1password
</code></pre>
