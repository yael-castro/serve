# Serve
![Tux, the Linux mascot](/docs/images/header.png)
CLI application that makes a simple http file server using one command. Made by [Yael](github.com/yael-castro)

###### Install

```
go install github.com/yael-castro/serve@latest
```

###### Usage

```
serve -port {port} -dir {directory}
```
###### Flags
<table>
    <tr>
        <th>CLI Flag</th>
        <th>Required value</th>
        <th>Default value</th>
    </tr>
    <tr>
        <td>port</td>
        <td>Integer</td>
        <td>8080</td>
    </tr>
        <tr>
        <td>dir</td>
        <td>String</td>
        <td>./</td>
    </tr>
</table>