/* src/router/Router.svelte generated by Svelte v3.23.0 */
import {
	SvelteComponent,
	check_outros,
	create_component,
	destroy_component,
	detach,
	element,
	group_outros,
	init,
	insert,
	mount_component,
	safe_not_equal,
	transition_in,
	transition_out
} from "/web_modules/svelte/internal.js";

import { onMount } from "/web_modules/svelte.js";
import { writable } from "/web_modules/svelte/store.js";
import { currentRoute } from "../store/index.js";

function create_fragment(ctx) {
	let section;
	let current;
	var switch_value = /*component*/ ctx[0];

	function switch_props(ctx) {
		return {};
	}

	if (switch_value) {
		var switch_instance = new switch_value(switch_props(ctx));
	}

	return {
		c() {
			section = element("section");
			if (switch_instance) create_component(switch_instance.$$.fragment);
		},
		m(target, anchor) {
			insert(target, section, anchor);

			if (switch_instance) {
				mount_component(switch_instance, section, null);
			}

			current = true;
		},
		p(ctx, [dirty]) {
			if (switch_value !== (switch_value = /*component*/ ctx[0])) {
				if (switch_instance) {
					group_outros();
					const old_component = switch_instance;

					transition_out(old_component.$$.fragment, 1, 0, () => {
						destroy_component(old_component, 1);
					});

					check_outros();
				}

				if (switch_value) {
					switch_instance = new switch_value(switch_props(ctx));
					create_component(switch_instance.$$.fragment);
					transition_in(switch_instance.$$.fragment, 1);
					mount_component(switch_instance, section, null);
				} else {
					switch_instance = null;
				}
			} else if (switch_value) {
				0;
			}
		},
		i(local) {
			if (current) return;
			if (switch_instance) transition_in(switch_instance.$$.fragment, local);
			current = true;
		},
		o(local) {
			if (switch_instance) transition_out(switch_instance.$$.fragment, local);
			current = false;
		},
		d(detaching) {
			if (detaching) detach(section);
			if (switch_instance) destroy_component(switch_instance);
		}
	};
}

function instance($$self, $$props, $$invalidate) {
	let { routes = [] } = $$props;
	const defaultComponent = { path: "/", component: null };
	let component;
	const storedRoutes = writable([]);

	// $: console.log('Routes reactive statement', $storedRoutes);    
	routes.forEach(route => {
		storedRoutes.update(r => [...r, route]);
		return route;
	});

	// setContext('addRoute', (route) => {
	// 	console.log('Route added');
	// 	storedRoutes.update(rs => {
	// 		rs.push(route);
	// 		return rs;
	// 	});
	// });
	onMount(() => {
		const currentRoute = routes.find(({ path }) => path === window.location.pathname);

		$$invalidate(0, component = currentRoute
		? currentRoute.component
		: defaultComponent.component);
	});

	currentRoute.subscribe(route => {
		const currentRoute = routes.find(({ path }) => path === route);

		$$invalidate(0, component = $$invalidate(0, component = currentRoute
		? currentRoute.component
		: defaultComponent.component));

		
	});

	$$self.$set = $$props => {
		if ("routes" in $$props) $$invalidate(1, routes = $$props.routes);
	};

	return [component, routes];
}

class Router extends SvelteComponent {
	constructor(options) {
		super();
		init(this, options, instance, create_fragment, safe_not_equal, { routes: 1 });
	}
}

export default Router;