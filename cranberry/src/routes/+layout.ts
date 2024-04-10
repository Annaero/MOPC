import { dev } from '$app/environment';
import { PUBLIC_SUPABASE_ANON_KEY, PUBLIC_SUPABASE_URL } from '$env/static/public'
import type { LayoutLoad } from './$types'
import { createBrowserClient, isBrowser, parse } from '@supabase/ssr'
import { inject } from '@vercel/analytics';

inject({ mode: dev ? 'development' : 'production' });

export const load: LayoutLoad = async ({ fetch, data, depends }) => {
    depends('supabase:auth')
    const session = data.session;

    const supabase = createBrowserClient(PUBLIC_SUPABASE_URL, PUBLIC_SUPABASE_ANON_KEY, {
        global: {
            fetch,
        },
        cookies: {
            get(key) {
                if (!isBrowser()) {
                    return JSON.stringify(session)
                }

                const cookie = parse(document.cookie)
                return cookie[key]
            },
        },
    })

    const profile = data.profile;

    let avatarUrl = null;
    if (profile) {
        try {
            const { data, error } = await supabase.storage
                .from("avatars")
                .download(profile.avatarUUID);

            if (error) {
                throw error;
            }

            avatarUrl = URL.createObjectURL(data);
        } catch (error) {
            if (error instanceof Error) {
                console.error("Error downloading image: ", error.message);
            }
        }
    }


    return { supabase, session, profile, avatarUrl }
}