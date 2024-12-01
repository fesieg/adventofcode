export async function getFileContent(filePath: string): Promise<string> {

  // Bun is a global object that is always available, but letting ts know that takes some work 
  // @ts-ignore
  const file = Bun.file(filePath);
  
  return file.text();
}