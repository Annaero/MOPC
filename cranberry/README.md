MOPC project.

# Enviroment variables

Create file `.env`

Specify `DATABASE_URL`, `DIRECT_URL`, `PUBLIC_SUPABASE_URL`, `PUBLIC_SUPABASE_ANON_KEY`, `PUBLIC_SUPABASE_URL`, `PUBLIC_SUPABASE_ANON_KEY`

For detaails see https://supabase.com/partners/integrations/prisma

# Generaty prisma clients and zod schema
```shell
npx prisma generate
```

# Update database with new schema
```shell
npx prisma migrate dev
```

~~# How prisma works with supabase auth~~
~~https://medium.com/@ngoctranfire/using-prisma-with-supabase-row-level-security-and-multi-schema-7c53418adba3~~