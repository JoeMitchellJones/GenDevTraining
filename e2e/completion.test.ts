import {
  commands,
  position,
} from './utils';

describe('Completion command', () => {
  jest.setTimeout(10000);
  it('Completion command', async () => {
    const testingFile = 'config1.yml';
    await commands.didOpen(testingFile);

    const res = await commands.complete(testingFile, position(1, 1));

    expect(res).toMatchSnapshot();
  });

  it('Complete list of jobs', async () => {
    const testingFile = 'invalid-files/autocomplete-jobs.yml';

    await commands.didOpen(testingFile);
    const res = await commands.complete(testingFile, position(85, 9));

    expect(res).toMatchSnapshot();
  });
});
