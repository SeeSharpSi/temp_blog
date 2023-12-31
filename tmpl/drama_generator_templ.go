// Code generated by templ@v0.2.364 DO NOT EDIT.

package templ

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func DramaGenerator() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		templBuffer, templIsBuffer := w.(*bytes.Buffer)
		if !templIsBuffer {
			templBuffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templBuffer)
		}
		ctx = templ.InitializeContext(ctx)
		var_1 := templ.GetChildren(ctx)
		if var_1 == nil {
			var_1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, err = templBuffer.WriteString("<div id=\"page\"><div class=\"container\"><div style=\"text-align: center;\"><p>")
		if err != nil {
			return err
		}
		var_2 := `Start drama with `
		_, err = templBuffer.WriteString(var_2)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("<span id=\"blank1\">")
		if err != nil {
			return err
		}
		var_3 := `_____`
		_, err = templBuffer.WriteString(var_3)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</span> ")
		if err != nil {
			return err
		}
		var_4 := `about `
		_, err = templBuffer.WriteString(var_4)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("<span id=\"blank2\">")
		if err != nil {
			return err
		}
		var_5 := `_____`
		_, err = templBuffer.WriteString(var_5)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</span> ")
		if err != nil {
			return err
		}
		var_6 := `and bring up `
		_, err = templBuffer.WriteString(var_6)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("<span id=\"blank3\">")
		if err != nil {
			return err
		}
		var_7 := `_____`
		_, err = templBuffer.WriteString(var_7)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</span></p><br><button id=\"generateButton\">")
		if err != nil {
			return err
		}
		var_8 := `Generate`
		_, err = templBuffer.WriteString(var_8)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</button></div></div></div><script>")
		if err != nil {
			return err
		}
		var_9 := `
        document.getElementById('generateButton').addEventListener('click', function() {
            const a = ["Laravel Users", "DevOps"];
            const b = ["Server-Side Rendering", "CLI vs GUI"];
            const c = ["'i use vim btw'", "how slow/fast Python is"];

            shuffleWord(a, 'blank1', function() {
                shuffleWord(b, 'blank2', function() {
                    shuffleWord(c, 'blank3');
                });
            });
        });

        function shuffleWord(array, elementId, callback) {
            const elem = document.getElementById(elementId);
            let counter = 0;

            const shuffleInterval = setInterval(function() {
                elem.textContent = array[Math.floor(Math.random() * array.length)];
                counter += 50; 

                if (counter >= 1200) { 
                    clearInterval(shuffleInterval);
                    elem.textContent = array[Math.floor(Math.random() * array.length)];
                    if (callback) callback();
                }
            }, 50);
        }
    `
		_, err = templBuffer.WriteString(var_9)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</script>")
		if err != nil {
			return err
		}
		if !templIsBuffer {
			_, err = templBuffer.WriteTo(w)
		}
		return err
	})
}
