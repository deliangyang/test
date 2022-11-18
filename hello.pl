
use warnings FATAL => 'all';
use strict;
use warnings;


while(<>) {
    if (/^use (.+)$/) {
        my $name = $1;
        print "$name\n"
    }

}

