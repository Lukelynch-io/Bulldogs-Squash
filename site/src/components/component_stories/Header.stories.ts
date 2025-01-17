import type { Meta, StoryObj } from '@storybook/vue3';
import Header from '../Header.vue';

const meta: Meta<typeof Header> = {
  component: Header,
};

export default meta;
type Story = StoryObj<typeof Header>;

/*
 *👇 Render functions are a framework specific feature to allow you control on how the component renders.
 * See https://storybook.js.org/docs/api/csf
 * to learn how to use render functions.
 */
export const Primary: Story = {
  render: (args) => ({
    components: { Header },
    setup() {
      return { args };
    },
    template: '<Header />',
  }),
  args: {
    isLoginShowing: () => true,
    loggedInUsername: "TestUser"
  },
};
