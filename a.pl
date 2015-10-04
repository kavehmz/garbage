$max=1000000;

sub is_prime {
	$n =shift;
	for (my $i = 2; $i <= int(sqrt(($n))); $i++){
		if ($n % $i == 0) {
			return;
		}
	}
	return 1;
}


for (my $i = 2; $i < $max; $i++) {
	if (is_prime($i)) {
		$x++
	}
}

print "$x\n";