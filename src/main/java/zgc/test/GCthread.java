package zgc.test;

import org.apache.commons.lang3.time.StopWatch;
import java.util.concurrent.TimeUnit;
import java.time.temporal.ChronoUnit;
import java.time.Duration;
import java.time.Instant;

public final class GCthread extends Thread {

	private static final class LatencyStats {
		static long min = 0;
		static long max = 0;
		static long mean = 0;
		static long count = 0;
	}

	private final static Instant time() {
		return Instant.now().truncatedTo(ChronoUnit.MICROS);
	}

	// helper function
	private final static void error(final Exception e) {
		System.out.println("Fatal Error: " + e);
		System.exit(1);
	}

	public final static void printLatencyStats() {
		LatencyStats.mean /= LatencyStats.count;
		System.out.println("Latency (min, max, mean): " +
				LatencyStats.min+" µs, " +
				LatencyStats.max+" µs, " +
				LatencyStats.mean+" µs"
		);
	}

	@Override
   public final void run ()
   {
		// wait on main thread to call System.gc()
		final var t0 = time();
		try {
			TimeUnit.MILLISECONDS.sleep(10);
		} catch (final InterruptedException e) {
			error(e);
		}

		// stats handling
		var latency = ((Duration.between(t0, time()).toNanos())/1000) - 10000;
		if (LatencyStats.min == 0 || latency < LatencyStats.min) {
			LatencyStats.min = latency;
		}
		LatencyStats.max = Math.max(LatencyStats.max, latency);
		LatencyStats.mean += latency;
		++LatencyStats.count;

		// output handling
		System.out.println("Latency: " + latency + " µs");
   }
}
