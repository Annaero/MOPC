import { browser } from '$app/environment'
import { init, register } from 'svelte-i18n'

const defaultLocale = 'en-GB'

register('en-GB', () => import('$lib/locales/en.json'))
register('ru-RU', () => import('$lib/locales/ru.json'))

init({
    fallbackLocale: defaultLocale,
    initialLocale: browser ? window.navigator.language : defaultLocale,
})