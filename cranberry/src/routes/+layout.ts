import { PUBLIC_SUPABASE_URL, PUBLIC_SUPABASE_KEY } from '$env/static/public'
import { createSupabaseLoadClient } from '@supabase/auth-helpers-sveltekit'
import type { Database } from '$lib/supabase'
import { dev } from '$app/environment';
import { inject } from '@vercel/analytics';
 
inject({ mode: dev ? 'development' : 'production' });

export const load = async ({ fetch, data, depends }) => {
    depends('supabase:auth')

    const supabase = createSupabaseLoadClient<Database>({
        supabaseUrl: PUBLIC_SUPABASE_URL,
        supabaseKey: PUBLIC_SUPABASE_KEY,
        event: { fetch },
        serverSession: data.session,
    })

    const {
        data: { session },
    } = await supabase.auth.getSession()

    return { supabase, session }
}