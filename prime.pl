our $num;
use Time::HiRes;
sub fill {
	my $i = shift;
	my $max = shift;
	my $a = 2 * $i;
	while ($a <= $max) {
		$nums->[$a] = 1;
		$a = $a + $i;
	}
}

my $max = 1000000;
my $m = int(sqrt($max));

$t0 = [Time::HiRes::gettimeofday];

for (my $i = 2; $i <= $m; $i = $i + 2) {

	if (!$nums->[$i]) {
		fill($i, $max);
	}
	if ($i == 2) {
		$i = 1;
	}

}

my $ps;
$c=0;
for (my $i = 2; $i <= $max; $i++){
	if (! $nums->[$i]) {
		$ps->[$c]=$i;
		$c++;
	}
}

$elapsed = Time::HiRes::tv_interval ( $t0 );


print scalar @{$ps}, "   [$elapsed] \n";

#use Data::Dumper;
#print Data::Dumper::Dumper($ps);