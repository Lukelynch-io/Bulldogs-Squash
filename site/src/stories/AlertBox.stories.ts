import type { Meta, StoryObj } from '@storybook/vue3'

import AlertBox from '../components/AlertBox.vue'
import { MessageType } from '@/datatypes/MessageType'


const meta: Meta<typeof AlertBox> = {
  component: AlertBox,
}

export default meta
type Story = StoryObj<typeof AlertBox>

/*
 *👇 Render functions are a framework specific feature to allow you control on how the component renders.
 * See https://storybook.js.org/docs/api/csf
 * to learn how to use render functions.
 */
export const Primary: Story = {
  render: (args) => ({
    components: { AlertBox },
    setup() {
      return { args }
    },
    template: '<AlertBox v-bind="args"/>',
  }),
  args: {
    messageType: MessageType.Info,
    message: "This is the message"
  },
}
