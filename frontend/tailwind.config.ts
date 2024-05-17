import type { Config } from 'tailwindcss'

const config: Config = {
  content: [
    './src/pages/**/*.{js,ts,jsx,tsx,mdx}',
    './src/components/**/*.{js,ts,jsx,tsx,mdx}',
    './src/app/**/*.{js,ts,jsx,tsx,mdx}',
  ],
  theme: {
      extend: {
      colors: {
		  'foodiee-black-primary': '#131515',
		  'foodiee-black-secondary': '#2B2C28',
		  'foodiee-white': '#FFFAFB',
		  'foodiee-button-primary': '#7DE2D1',
		  'foodiee-button-secondary': '#339989',
      },
      backgroundImage: {
		  'gradient-radial': 'radial-gradient(var(--tw-gradient-stops))',
		  'gradient-conic':
		  'conic-gradient(from 180deg at 50% 50%, var(--tw-gradient-stops))',
      },
    },
  },
  plugins: [],
}
export default config
