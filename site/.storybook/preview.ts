import type { Preview, Decorator } from '@storybook/vue3'

import '../style.css'
import '../colours.css'
export const globalTypes = {
  theme: {
    name: 'Dark Mode',
    description: 'Enable dark mode',
    defaultValue: 'light', // Default to 'light' mode
    toolbar: {
      icon: 'circlehollow', // Icon to represent the toggle
      items: [
        { value: 'light', title: 'Light Mode' },
        { value: 'dark', title: 'Dark Mode' },
      ],
      showName: true,
    },
  },
};

const preview: Preview = {
  decorators: [
    (_, context) => {
      const theme = context.globals.theme || 'light'
      var element = document.getElementsByTagName("body")[0];
      if (theme == 'light') {
        element.classList.remove('dark');
      } else {
        element.classList.add('dark');
      }
      return {
        template: '<story/>',
      };
    }
  ],
  parameters: {
    controls: {
      matchers: {
        color: /(background|color)$/i,
        date: /Date$/i,
      },
    },
  },
}

export default preview
