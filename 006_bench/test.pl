use Time::HiRes qw( gettimeofday tv_interval 
		       );
use Contract;

$a=Contract->new();

$t0 = [gettimeofday];
$n=0;
for (my $i = 0; $i < 10000000; $i++) {
	$n++;
}
print "loop time :", tv_interval ( $t0 ), "\n";

$t0 = [gettimeofday];
$n=0;
for (my $i = 0; $i < 10000000; $i++) {
	$n=exp($i);
}
print "exp time :", tv_interval ( $t0 ), "\n";


$t0 = [gettimeofday];
for (my $i = 0; $i < 10000; $i++) {
	@a=();
	#print $a[10];
	for (my $j = 0; $j < 10000; $j++) {
		push @a,$j;
	}
}
print "append time :", tv_interval ( $t0 ), "\n";

$t0 = [gettimeofday];

for (my $i = 0; $i < 10000000; $i++) {
	$p=Contract->new();
}
print "package time :", tv_interval ( $t0 ), "\n";
