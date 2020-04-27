package zgc.test;

import java.lang.management.GarbageCollectorMXBean;
import java.lang.management.ManagementFactory;
import java.util.List;

public class GCinfo {
	public static final void print() {
		try {
			System.out.println("-".repeat(10)+" START GC INFO "+"-".repeat(10));
			final List<GarbageCollectorMXBean> gcMxBeans = ManagementFactory.getGarbageCollectorMXBeans();
			for (final GarbageCollectorMXBean gcMxBean : gcMxBeans) {
				System.out.println(gcMxBean.getName());
				System.out.println(gcMxBean.getObjectName());
			}
			System.out.println("-".repeat(10)+" END GC INFO "+"-".repeat(10));
		} catch (final RuntimeException re) {
			throw re;
		} catch (final Exception exp) {
			throw new RuntimeException(exp);
		}
   }
}
