import type { Meta, StoryObj } from '@storybook/vue3';

import BlogPost from '../BlogPost.vue';

const meta: Meta<typeof BlogPost> = {
  component: BlogPost,
};

export default meta;
type Story = StoryObj<typeof BlogPost>;

/*
 *👇 Render functions are a framework specific feature to allow you control on how the component renders.
 * See https://storybook.js.org/docs/api/csf
 * to learn how to use render functions.
 */
export const Primary: Story = {
  render: (args, { globals }) => ({
    components: { BlogPost },
    setup() {
      return { args, theme: globals.theme };
    },
    template: '<BlogPost v-bind="args"/>',
  }),
  args: {
    title: "Bulldogs Competition",
    description: "Test Description",
    imageUrl: "ImageUrl",
  }
};
