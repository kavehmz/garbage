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
for (my $i = 0.0; $i < 10000000.0; $i++) {
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


my $SQRT2PI = 2.506628274631;
sub pdf {
  my ( $x, $m, $s ) = ( 0, 0, 1 );
  $x = shift if @_;
  $m = shift if @_;
  $s = shift if @_;

  if( $s <= 0 ) {
    die "die";
  }

  my $z = ($x-$m)/$s;

  return exp(-0.5*$z*$z)/($SQRT2PI*$s);
}
$t0 = [gettimeofday];
for (my $i = 0; $i < 1000000; $i++) {
	$p=pdf(1.2,1.3,1.4);
}
print "pdf time :", tv_interval ( $t0 ), "\n";
