import type { Meta, StoryObj } from '@storybook/vue3'

import LoginForm from '../LoginForm.vue'

const meta: Meta<typeof LoginForm> = {
  component: LoginForm,
}

export default meta
type Story = StoryObj<typeof LoginForm>

/*
 *👇 Render functions are a framework specific feature to allow you control on how the component renders.
 * See https://storybook.js.org/docs/api/csf
 * to learn how to use render functions.
 */
export const Primary: Story = {
  render: (args) => ({
    components: { LoginForm },
    setup() {
      return { args }
    },
    template: '<LoginForm v-bind="args"/>',
  }),
  args: {
  },
}
