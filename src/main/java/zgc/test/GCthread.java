package zgc.test;

import org.apache.commons.lang3.time.StopWatch;
import java.util.concurrent.TimeUnit;
import java.time.temporal.ChronoUnit;
import java.time.Duration;
import java.time.Instant;

public final class GCthread extends Thread {

	private final static Instant time() {
		return Instant.now().truncatedTo(ChronoUnit.MICROS);
	}

	// helper function
	private final static void error(final Exception e) {
		System.out.println("Fatal Error: " + e);
		System.exit(0);
	}

	@Override
   public final void run ()
   {
		final var t0 = time();
		try {
			TimeUnit.MILLISECONDS.sleep(10);
		} catch (final InterruptedException e) {
			error(e);
		}
		System.out.println("Latency: " + (((Duration.between(t0, time()).toNanos())/1000) - 10000) + " µs");
   }
}
