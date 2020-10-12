{ pkgs ? import <nixpkgs> {} }:
pkgs.stdenv.mkDerivation {
	name = "ClassroomBot-devenv";
	buildInputs = with pkgs; [
		go
	];
	shellHook = ''
		export GOPATH=$HOME/go
		source .env 2>/dev/null >/dev/null
	'';
}
