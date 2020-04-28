package zgc.test;

import java.util.LinkedHashSet;
import java.util.Iterator;
import java.util.Set;

public final class lru_cache<T> {

	Set<T> cache;
	int capacity;

	public static final void test() {
		final var c = new lru_cache<Integer>(4);
		c.refer(1);
		c.refer(2);
		c.refer(3);
		c.refer(1);
		c.refer(4);
		c.refer(5);
		c.display();
	}

	public lru_cache(final int _capacity) {
		cache = new LinkedHashSet<T>(_capacity);
		capacity = _capacity;
	}

	public final void refer(final T key) {
		if (no(key)) {
			emplace(key);
		}
	}

	private final boolean no(final T key) {
		if (cache.contains(key)) {
			cache.remove(key);
			cache.add(key);
			return false;
		} else {
			return true;
		}
	}

	private final void display() {
		final var itr = cache.iterator();
		while (itr.hasNext()) {
			System.out.println(itr.next());
		}
	}

	private final void emplace(final T key) {
		if (cache.contains(key)) {
			cache.remove(key);
		} else if (cache.size() == capacity) {
			cache.remove(cache.iterator().next());
		}
		cache.add(key);
	}
}
