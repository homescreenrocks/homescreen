import { HomescreenPage } from './app.po';

describe('homescreen App', () => {
  let page: HomescreenPage;

  beforeEach(() => {
    page = new HomescreenPage();
  });

  it('should display message saying app works', () => {
    page.navigateTo();
    expect(page.getParagraphText()).toEqual('hs works!');
  });
});
