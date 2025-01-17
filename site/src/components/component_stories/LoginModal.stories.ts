import type { Meta, StoryObj } from '@storybook/vue3'

import LoginModal from '../LoginModal.vue'

const meta: Meta<typeof LoginModal> = {
  component: LoginModal,
}

export default meta
type Story = StoryObj<typeof LoginModal>

/*
 *👇 Render functions are a framework specific feature to allow you control on how the component renders.
 * See https://storybook.js.org/docs/api/csf
 * to learn how to use render functions.
 */
export const Primary: Story = {
  render: (args) => ({
    components: { LoginModal },
    setup() {
      return { args }
    },
    template: '<LoginModal v-bind="args"/>',
  }),
  args: {
    flag: true,
  },
}
