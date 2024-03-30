import { supabase } from '$lib/supabase/client.js'
import type { Database } from '$lib/supabase'
import { dev } from '$app/environment';
import { inject } from '@vercel/analytics';
import type { SupabaseClient } from '@supabase/supabase-js';

inject({ mode: dev ? 'development' : 'production' });

export const load = async ({ fetch, data, depends }) => {
    // depends('supabase:auth')

    const {
        data: { session },
    } = await supabase.auth.getSession()

    return { supabase, session }
}