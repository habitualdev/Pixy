<h1 style="text-align: center">Pixy</h1>
<h3 style="text-align:center"> All-in-one DHCP, TFTP, and HTTP PXE server  </h3>

<div>
All in one solution for the beast of a problem that PXE usually offers up. 
</div>
<hr>
<h3 style="text-align: center">TFTP</h3>
<div>
All pxeboot image files are embedded inside the binary for ease of use. Designed to use pxelinux as the bootloader. 
The "default" file in the root directory is served as "pxelinux.cfg" allowing users to modify it even at runtime.
</div>
<h3 style="text-align: center">HTTP</h3>
<div>
Intended to be the source of all boot images.
</div>
<h3 style="text-align: center">DHCP</h3>
<div>
Gives out addresses using a golang DHCP implementation. 
</div>
<hr>
<h3 style="text-align: center">Usage</h3>

Running the binary without arguments:
 - Attempts to locate the default network interface and IP address, and binds to those

With arguments:
 - ./Pixy \<interface> \<ip address>
 - Attempts to bind to supplied interface and address

Requires sudo/root on linux

